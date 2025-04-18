# DSC First Resource Tutorial

Code created going through the sample tutorial https://powershell.github.io/DSC-Samples/tutorials/first-resource/

# Notes
- --inputJSON fails to parse with "Error: invalid argument "{scope:machine}" for "--inputJSON" flag: invalid character 's' looking for beginning of object key string" - this happens on the example code as well https://github.com/PowerShell/DSC-Samples/tree/main/samples/go/resources/first
- Had to specify "./tstoy" to make the executable run, took a bit of debugging as the error was otherwise just "Exit status code 1", would probably not be a problem on linux
- schema is off and `dsc resource list` fails to parse it because of "manifestVersion" entry (and missing $schema), I just removed the one and added the other from https://learn.microsoft.com/en-us/powershell/dsc/reference/schemas/resource/manifest/root?view=dsc-3.0 and everything just worked.