Following along https://interpreterbook.com and making changes/simplification/cleanups

Install/run it:
```shell
CGO_ENABLED=0 go install -trimpath -ldflags="-w -s" -tags no_net,no_json github.com/ldemailly/gorepl@latest
```

Sample:
```shell
gorepl -parse
$ fact = func(n) {if (n<=1) {return 1} n*fact(n-1)}
$ n=fact(6)
== Parse ==> (n = fact(6))
== Eval  ==> 720
$ m=fact(7)
== Parse ==> (m = fact(7))
== Eval  ==> 5040
$ m/n
== Parse ==> (m / n)
== Eval  ==> 7
```

See also [sample.gr](sample.gr) that you can run with
```
gorepl *.gr
```

Dev mode:
```shell
go install golang.org/x/tools/cmd/stringer@latest
make # for stripped down executable including build tags etc to make it minimal
```

Status: All done: ie functional int, string and boolean expressions, functions, lambdas, arrays, maps,
print, log, and more


### Reading notes

- See the commit history for improvements/changes (e.g redundant state in lexer etc)

- [x] interface nil check in parser

- [x] Do we really need all these `let `, wouldn't `x = a + 3` be enough? made optional

- [ ] Seems like ast and object are redundant to a large extent

- [x] Introduced errors sooner, it's sort of obviously needed

- [x] Put handling of return/error once at the top instead of peppered all over

- [x] Make all the Eval functions receiver methods on State instead of passing environment around

- [x] made built ins like len() tokens (cheaper than carrying the string version during eval)

- [ ] fix up == and != in 3 places (int, string and default)

- [ ] change int to ... float? number? or rather add float/double (maybe also or big int?...)

- [ ] use + for concat of arrays and merging of maps

- [x] call maps maps and not hash (or maybe assoc array but that's long)

- [x] don't make a slice to join with , when there is already a strings builder. replace byte buffers by string builder.

- [x] generalized tokenized built in (token id based instead of string)

- [ ] Add "extension" internal functions (calling into a go function), with variadic params, param types etc

- [x] Identifiers are letter followed by alphanum*

- [x] map of interface correctly equals the actual underlying types, no need for custom hashing
  -> implies death to pointers (need non pointer receiver and use plain objects and not references)

- [ ] unicode

- [x] flags for showing parse or not (default not pass `-parse` to see parsing)

- [x] file input vs stdin repl (made up .gr for gorepl)

- [ ] actual name for the language - it's not monkey (though it's monkey compatible, just better/simpler/...)

- [ ] multiline support in stdin repl

- [x] add >= and <= comparaison operators

- [x] add comments support (line)
   - [ ] add /* */ style

- [ ] line numbers for errors (for file mode)

- [x] use `func` instead of `fn` for functions

- [x] figure out how to get syntax highlighting (go style closest - done thx to viulisti -> .gitattributes)

- [ ] assignment to maps and arrays

- [ ] for loop

### Usage

```
gorepl 0.15.0 usage:
	gorepl [flags] *.gr files to interpret or no arg for stdin repl...
or 1 of the special arguments
	gorepl {help|envhelp|version|buildinfo}
flags:
  -eval
    	show eval results (default true)
  -parse
    	show parse tree
  -shared-state
    	All files share same interpreter state (default is new state for each)
```
(excluding logger control, see `gorepl help` for all the flags, of note `-logger-no-color` will turn off colors for gorepl too)
