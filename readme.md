<kbd> <img src="https://github.com/ManuCiao10/eagle/blob/master/handler/mods/git.png" /> </kbd>

### Features

- [x] Bot Auto-Update + Loader
- [x] AWS API Auth
- [x] Auth API Discord
- [x] Discord Rich Presence
- [x] Logger
- [x] TLS client
- [x] Console Windows
- [x] Proxy handler
- [x] Webhook + Dashboard
- [x] Modules Manager
- [x] Task manager
- [x] Clean Architecture

### Build
Windows:
```
go generate
env GOOS=windows GOARCH=amd64 go build -o EagleBot_<version>.exe github.com/eagle
```

macOs(64-bit):
```
go generate
rm resource.syso
env GOOS=darwin GOARCH=amd64 go build -o EagleBot_<version> github.com/eagle
```

### TODO
- Rewrite the authentication for WHOP (Hyper has been deprecated)

### Download and Install
1.  From the dashboard install the executable file.exe
2.  Create a Folder Named EagleBot and insert the Executable 
3.  Run the Bot with a double click on the executable

### Artchitecture
Each of these [architectures](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html) produce systems that are:
1. Independent of Frameworks. The architecture does not depend on the existence of some library of feature laden software. This allows you to use such frameworks as tools, rather than having to cram your system into their limited constraints.
2. Testable. The business rules can be tested without the UI, Database, Web Server, or any other external element.
3. Independent of UI. The UI can change easily, without changing the rest of the system. A Web UI could be replaced with a console UI, for example, without changing the business rules.
4. Independent of Database. You can swap out Oracle or SQL Server, for Mongo, BigTable, CouchDB, or something else. Your business rules are not bound to the database.
5. Independent of any external agency. In fact your business rules simply don’t know anything at all about the outside world.

### Contact
-Discord: manuciao_dev
