<p align="center">
  <img src="assets/kiwi-logo.png" alt="Kiwi KV Store logo" width="256" />
</p>

<h1 align="center">kiwi — a tiny key–value storage</h1>

<p align="center">
  <b>Status: Work in Progress — not production-ready</b>
</p>

---

## What is kiwi?

Kiwi is a simple, educational key–value store project in Go. The goal is to make KV-store fundamentals tangible: in-memory index, append-only log, a small API, and crash recovery on startup.

> Why “Kiwi”?
> A playful nod to “KV” (key–value). Spoken in English (“kay-vee”), it sounds a bit like the german fruit “Kiwi” — and thus a friendly fruit mascot was born.

---

## Features (planned)

- In-memory key–value storage
- Append-only active log
- Crash recovery on startup (rebuild from log)
- Minimal API:
    - put
    - get
    - delete

---

## Roadmap

- v0.1
    - [ ] In-memory KV store
    - [ ] put/get/delete
    - [ ] Append-only active log and crash recovery
    - [ ] Simple API

Future ideas (nice-to-have):
- [ ] Persistence/checkpoints
- [ ] Concurrency/locking
- [ ] Metrics/monitoring

---

## Contributing

Contributions are welcome! Please:
- Open an issue to discuss changes
- Use clear commit messages
- Run go fmt / lint before committing

---

## License

MIT