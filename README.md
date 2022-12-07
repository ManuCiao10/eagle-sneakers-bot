<kbd> <img src="https://github.com/ManuCiao10/eagle/blob/master/handler/mods/git.png" /> </kbd>

# Structure

1. The architecture does not depend on the existence of some library of feature laden software.
2. Testable. The business rules can be tested without the UI, Database, Web Server, or any other external element.
   https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html

#OTHERS
 MODULES
 MONITOR
 MQT

#GUIDE
 DASHBOARD

#SECURITY
 1. https://stackify.com/what-is-inetpub/
 2. Sniffer tipo proxyman, fiddler, burp

#DOWNLOAD USER
 1. From the dashboard install the executable file.exe
 2. Create a Folder Named EagleBot and insert the Executable
 3. Run the Bot with a double click
 4. You will have the Auto-Update features

#PUSH UPDATE
 1. Run: env GOOS=windows GOARCH=amd64 go build -o EagleBot_<version>.exe github.com/eagle
 3. Upload the file in the dashboard 