# Ngulik NATS.io

Contoh contoh kecil implementasi dari NATS.io yang ditulis di [SOTOY](https://sotoy.onrender.com/posts/natsio/ngulik-natsio.html).

## Jalanin NATS

```bash
docker run --rm -p 4222:4222 -p 8222:8222 -p 6222:6222 --name nats-server -ti nats:latest
```