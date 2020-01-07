# Overview
This library contains a entity framework for bbolt, including CRUD (Create, Read, Update, Delete) and query operations, a filtering DSL (Domain Specific Language) as well as an AST (Abstract Syntax Tree) for the DSL

  * `zitiql` contains the ANTLR4 based grammar for the filtering DSL
      * Use the `./generate.sh` script to generate the go code when the grammar changes
  * `ast` contains an AST which can be built from the grammar. 
      * It is structured to be able to run filters across some set of data as well as interface for getting typing information about the symbols in the underlying datastore
      * It includes a visitor pattern which can be used for AST validations or transformations
      * It includes type transformation code, allowing the AST to be transformed into a typed version of itself
  * `boltz` contains:
       * Entity management framework, including entity store interfaces and base types for doing CRUD and querying
       * Code for managing bbolt indexes and foreign keys
       * bbolt utilities for getting/setting typed data from bbolt (see typed_bucket.go) 
       * Implementation of AST interfaces for reading symbols and getting symbol type information

# ANTLR patches
## Race condition
Note: The ANTLR4 Go implementation has a race condition, detected by go -race. 
The issue is tracked here: https://github.com/antlr/antlr4/issues/2040 
The issue was fixed as described here: https://github.com/google/cel-go/pull/177/files
It's not clear that this fix is required, but it appeases the race condition detector and it seems safer to have it than not.
The fix should be reapplied after regenerating source. The next time source is regenerated, see if we can automate this
in the generate.sh script by including the patch.
   
## ARM 32 Patch
A compile issue happens on arm 32 which required the following patch:

```diff
diff --git a/storage/zitiql/zitiql_parser.go b/storage/zitiql/zitiql_parser.go
index af80e9d..8732ec4 100644
--- a/storage/zitiql/zitiql_parser.go
+++ b/storage/zitiql/zitiql_parser.go
@@ -1170,7 +1170,7 @@ func (p *ZitiQlParser) Start() (localctx IStartContext) {
        p.GetErrorHandler().Sync(p)
        _la = p.GetTokenStream().LA(1)
 
-       for ((_la-3)&-(0x1f+1)) == 0 && ((1<<uint((_la-3)))&((1<<(ZitiQlParserLPAREN-3))|(1<<(ZitiQlParserBOOL-3))|(1<<(ZitiQlParserALL_OF-3))|(1<<(ZitiQlParserANY_OF-3))|(1<<(ZitiQlParserCOUNT-3))|(1<<(ZitiQlParserISEMPTY-3))|(1<<(ZitiQlParserNOT-3))|(1<<(ZitiQlParserIDENTIFIER-3)))) != 0 {
+       for ((_la-3)&-(0x1f+1)) == 0 && ((uint64(1)<<uint((_la-3)))&((1<<(ZitiQlParserLPAREN-3))|(1<<(ZitiQlParserBOOL-3))|(1<<(ZitiQlParserALL_OF-3))|(1<<(ZitiQlParserANY_OF-3))|(1<<(ZitiQlParserCOUNT-3))|(1<<(ZitiQlParserISEMPTY-3))|(1<<(ZitiQlParserNOT-3))|(1<<(ZitiQlParserIDENTIFIER-3)))) != 0 {
                {
                        p.SetState(141)
                        p.Query()
```

It may be a boundary condition and may be fixed if we add another constant, or it may not. Either way, the fix may need to be reapplied when the grammer is regenerated.