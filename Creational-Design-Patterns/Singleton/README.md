#### Singletons can be created by:
1. mutexes.
   - This is version 1 of the code.
2. sync.once.
   - This is version 2 of the code.
3. init() function:
   - We can create a single instance inside the init function.
   - This is only applicable if the early initialization of the object is possible.
   - The init function is only called once per file in a package.
   - So we can be sure that only a single instance will be created.
