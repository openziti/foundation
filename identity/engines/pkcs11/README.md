PKCS#11 engine
==============

# Testing
Tests in this package use [Software HSM](https://github.com/opendnssec/SoftHSMv2) for testing. 
The tests will skip if softhsm2 driver is not found. Install `softhsm2` if you want to run these tests.
Directory `softhsm-testdata` contains predefined token with keys expected by the tests.