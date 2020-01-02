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
