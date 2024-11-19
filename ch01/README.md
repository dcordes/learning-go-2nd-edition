# Chapter 01

## Installing the Go Tools / Your First Go Program

* for a new program
  * `go mod init hello_world`
* some maybe duplicates notes from `https://go.dev/doc/tutorial/getting-started`
  * dependency management
    * your code imports "packages" provided by "modules"
    * those dependencies are managed by a "go.mod" file
    * the file is created by issuing a `go mod init` command and giving it the name of the module your code will be in; that name is the module's module path
    * typically the file is in the root directory of the repository containing your module
    * if you issued this command from the directory `hello` it would be valid
      * `go mod init example/hello`
