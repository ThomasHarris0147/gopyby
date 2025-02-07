1. gem install quartz https://github.com/DavidHuie/quartz - utilizes the sharing of code between the 2 languages using net/rpc
2. create go code, everything has to be in structs
3. inside the main function you need to register each function (see adder.go).
4. inside ruby file you need to require quarts.
5. start a new client and import that go file and after that you can directly use those go file's functions
6. `ruby quartz_example.rb`