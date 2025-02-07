from wasmtime import Store, Module, Instance, Engine, Linker

# Create a Wasmtime engine and store
engine = Engine()
store = Store(engine)

# Load the WebAssembly module from file
module = Module.from_file(engine, "./add.wasm")

# Create a linker and instantiate the module
linker = Linker(engine)
instance = linker.instantiate(store, module)

# Call the `add` function with two numbers
add_function = instance.exports(store)["add"]
result = add_function(store, 5, 7)  # Change numbers as needed

print("Result:", result)  # Should print: 12
