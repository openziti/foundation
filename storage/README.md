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

Note: The ANTLR4 Go implementation has a race condition, detected by go -race. 
The issue is tracked here: https://github.com/antlr/antlr4/issues/2040 
The issue was fixed as described here: https://github.com/google/cel-go/pull/177/files
It's not clear that this fix is required, but it appeases the race condition detector and it seems safer to have it than not.
The fix should be reapplied after regenerating source. The next time source is regenerated, see if we can automate this
in the generate.sh script by including the patch.
   
