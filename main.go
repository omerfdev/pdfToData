package main

import (
    "encoding/json"
    "log"
    "net/http"
    "os"

    pdfcontent "github.com/unidoc/unipdf/v3/contentstream"
    "github.com/unidoc/unipdf/v3/extractor"
    "github.com/unidoc/unipdf/v3/model"
)

func main() {
    http.HandleFunc("/api/pdfdata", func(w http.ResponseWriter, r *http.Request) {
        // POST isteğini işliyoruz.
        if r.Method != http.MethodPost {
            w.WriteHeader(http.StatusMethodNotAllowed)
            return
        }

        // PDF dosyasının yolu form verisinden alınıyor.
        pdfDosyaYolu := r.FormValue("pdfdosyayolu")
        if pdfDosyaYolu == "" {
            http.Error(w, "PDF dosyası yolu belirtilmedi", http.StatusBadRequest)
            return
        }

        // PDF dosyasını açıyoruz.
        pdfFile, err := os.Open(pdfDosyaYolu)
        if err != nil {
            http.Error(w, "PDF dosyasını açarken hata oluştu", http.StatusInternalServerError)
            return
        }
        defer pdfFile.Close()

        // PDF okuyucuyu oluşturuyoruz.
        pdfReader, err := model.NewPdfReader(pdfFile)
        if err != nil {
            http.Error(w, "PDF okuyucu oluşturulurken hata oluştu", http.StatusInternalServerError)
            return
        }

        // PDF sayfa sayısını alıyoruz.
        numPages, err := pdfReader.GetNumPages()
        if err != nil {
            http.Error(w, "PDF sayfa sayısı alınırken hata oluştu", http.StatusInternalServerError)
            return
        }

        // JSON verilerini tutacak bir map oluşturuyoruz.
        jsonVerileri := make(map[string]string)

        // Her sayfa için içerikleri alıyoruz.
        for pageNum := 1; pageNum <= numPages; pageNum++ {
            // Sayfa içeriğini alıyoruz.
            page, err := pdfReader.GetPage(pageNum)
            if err != nil {
                http.Error(w, "PDF sayfası alınırken hata oluştu", http.StatusInternalServerError)
                return
            }

            // Sayfa içeriğini çıkartıyoruz.
            extractor, err := extractor.New(page)
            if err != nil {
                http.Error(w, "PDF sayfa çıkartıcısı oluşturulurken hata oluştu", http.StatusInternalServerError)
                return
            }

            // Metin öğelerini alıyoruz.
            extractedText, err := extractor.ExtractText()
            if err != nil {
                http.Error(w, "PDF metinleri çıkartılırken hata oluştu", http.StatusInternalServerError)
                return
            }

            // PDF sayfasındaki metinleri işliyoruz.
            processText(extractedText, jsonVerileri)
        }

        // JSON verilerini cevap olarak gönderiyoruz.
        jsonData, err := json.Marshal(jsonVerileri)
        if err != nil {
            http.Error(w, "JSON verisi oluşturulurken hata oluştu", http.StatusInternalServerError)
            return
        }
        w.Header().Set("Content-Type", "application/json")
        w.Write(jsonData)
    })

    // Sunucuyu başlatıyoruz.
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func processText(text string, jsonVerileri map[string]string) {
    // PDF sayfasındaki metinleri işlemek için bir fonksiyon.
    // Burada istediğiniz düzenlemeleri yapabilirsiniz.

    // Örneğin, metinde "Name Surname" veya "Ad Soyad" geçiyorsa,
    // onun hemen ardından gelen metni alıp JSON verilerine ekleyebiliriz.
    if idx := pdfcontent.StrFind(text, "Name Surname"); idx != -1 {
        value := getTextAfter(text, idx)
        jsonVerileri["NameSurname"] = value
    }
    if idx := pdfcontent.StrFind(text, "Ad Soyad"); idx != -1 {
        value := getTextAfter(text, idx)
        jsonVerileri["AdSoyad"] = value
    }
    // Benzer şekilde "Tel" geçiyorsa, onun ardından gelen metni alıp JSON verilerine ekleyebiliriz.
    if idx := pdfcontent.StrFind(text, "Tel"); idx != -1 {
        value := getTextAfter(text, idx)
        jsonVerileri["Tel"] = value
    }
}

func getTextAfter(text string, idx int) string {
    // Verilen dizinden sonraki metni alır.
    text = text[idx:]
    // Satır sonlarına göre bölerek ilk satırı alırız.
    lines := pdfcontent.SplitText(text)
    if len(lines) > 0 {
        return lines[0]
    }
    return ""
}
