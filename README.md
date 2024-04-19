# PDF Veri Çıkartma Uygulaması

Bu uygulama, PDF dosyalarından metin verisi çıkartarak belirli metin öğelerini işleyip JSON formatında sunar. Özellikle belirli metin öğelerinin bulunduğu PDF dosyalarından bilgi toplamak için kullanılabilir.

## Kurulum

1. Bu kodu çalıştırmak için öncelikle Go programlama dilinin yüklü olması gerekmektedir.
2. Uygulama dosyalarını bilgisayarınıza indirin veya klonlayın.
3. Terminal veya komut istemcisinde projenin bulunduğu dizine gidin.
4. `go run main.go` komutunu kullanarak uygulamayı başlatın.
5. Tarayıcınızdan `http://localhost:8080/api/pdfdata` adresine giderek PDF dosyalarından çıkartılan verileri alın.

## Kullanım

- `/api/pdfdata` endpoint'i üzerinden PDF dosyalarından çıkartılan verileri alabilirsiniz. 
- İsteğinizi POST metodu ile yapmalısınız ve `pdfdosyayolu` parametresi ile PDF dosyasının yolunu belirtmelisiniz.
- Uygulama, PDF dosyasının her bir sayfasını tarar, belirli metin öğelerini bulur ve bunları JSON formatında sunar.

## Lisans

Bu proje MIT lisansı altında lisanslanmıştır. Daha fazla bilgi için `LICENSE` dosyasını inceleyebilirsiniz.
