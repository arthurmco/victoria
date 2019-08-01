# victoria

An (unofficial) language server to the [V programming language](https://vlang.io/).

To know more about what a language server is, check out [the LSP
official site](https://langserver.org/)

## Development status
I am very early on development. _Very_ early, like a baby that has
just been born. But you still can help, giving your opinion, or in
code.

Setup instructions will appear soon, when the language server start
working.

## Planned features
 - **Code completion**  
   Means that, given a type, the lang server will show all available
   methods for that type. Lexical context must be respected: inside
   the method, all methods need to be shown, and outside of it, only
   public methods

 - **Hover**  
   When we hover a variable, show the type (?) and places where it
   appears on the file, For example, if you hover over an variable
   called 'foo', it will highlight the location of all uses of said
   variable in the current file.

 - **Jump to definition**  
   Given a function, jump to where the function is defined.

 - **Workspace symbols** (Implemented after 1.0)  
   Show every symbol (function, variable, module) used by this code
 
 - **Block folding** (Implemented after 1.0)  
   Tells the editor that a range of code is a block that can be
   folded. Useful, but not too urgent.
 
 - **Diagnostices**  
   Show errors in your code as you type.  
   This might not be supported, at least for now, because it does not
   uses the compiler, and: 
    - compiling everytime would be overkill.
    - copying the compiler messages would be unsafe, as they could
      change
   BUUUT, if there is enough demand, this can be implemented

I am planning to focus on code completion and jump to definition,
since it is what I use more on other languages, but you can also give
your opinion on this.

## Questions
### Why Go? Why not V?
Because the language is on alpha. It is easier to change the parser to
recognize a change in the language syntax than change the parser *and*
the language.
But don't worry, I will rewrite it in V when the language reaches
stable.

### Why not `vls` or something like this?
I am not the author of V, nor affiliated with him. I do not even know
him. Therefore, I didn't want to stole this obvious name from him, in
the case he want to create a language server for his language. 
