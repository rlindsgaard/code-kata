# On Duty kata

The duty rotation generates the data file. The on duty reads the file and determines who is actually on call


## Initial prompt

Generate a powershell cmdlet that outputs who is on duty. 

It should take the path of a powershell datafile as script parameter that contains the duty roster specified as a list of hashtables with Date and Name keys. Date is a string formatted by yyyy-mm-dd

The person on duty is determined by what Date has last passed given by a  script parameter that defaults to the current date.

The answer should include comment based help content, conform to powershell naming conventions and adhere to clean code design practices.

Explain your design

#### Notes

- The more domain context I provide, the better the help section and variable names. This is a time-saver!
- Separating into natural language paragraphs helps distinguish the various requirements from one another especially as they get more complex. I’ll make a note of that (here it was context, param, calculation, param, additonal requirements) 
- “clean code design practices” did not do anything noticeable to the code structure it is, I reckon, too generic a formulation and too debatable anyhow. 
- “explain your design” yielded a re-formulation of the code and structure it is not what I was hoping for but I did not know exactly what to expect either. I like the idea of a “gotcha, I will do…” recipient re-formulation of a task to help identify any misunderstandings.

- “Sorts the filtered entries in descending order to find the most recent date.”
  ```
  $lastDutyEntry = $DutyRoster |
        Where-Object { $_.Date -le $currentDateString } |
        Sort-Object { [datetime]::ParseExact($_.Date, ‘yyyy-MM-dd’, $null) } -Descending |
        Select-Object -First 1
  ``` 
  this is really great! Here I think the machine proves more powerfull than a human as we as humans are biased to expect order in the input because that we saw as example (at least until lesson learned and proven otherwise). And I personally fall into this trap once in a while.
  It is also a note to self that I need to be explicit about the assumptions we can make towards the i input or domain.
  
### Refinement 1

regenerate your answer constraining the powershell datafile to be a .psd1 file

#### Notes
- “powershell datafile” did not have the intended result, specifying the file-type did.

## Unit tests

Now generate a pester file that, assuming it is placed in the same directory, tests the Get-OnDutyPerson cmdlet. Make sure to include tests covering all code paths of the cmdlet. Make sure that tests will succeed indepent of the actual current date.

#### Notes

- “I’m sorry but there was an error. Please try again.” that was neither intended nor expected. Trying again.

### Refinement 1
Now generate a pester file that, assuming it is placed in the same directory, tests the Get-OnDutyPerson cmdlet. Make sure to include tests covering all code paths of the cmdlet. Use mocking to ensure that tests will succeed indepent of the date the test executes.

#### Notes
- Yep, I needed to help hint how to accomplish that difficult task. Note to self: tell it when to mock

- “Make sure to include tests covering all code paths of the cmdlet” worked well and gave me what I expected, this is a good formulation I will take with me.

### Refinement 2

regenerate the test script such that files and folders created during tests are temporary

#### Notes

```
AfterAll {
        # Remove the temporary directory and its contents
        Remove-Item -LiteralPath $tempDir.FullName -Recurse -Force
    }
```

- I got what I asked for but we can do better. 

### Refinement 3

regenerate the test script such that file system operations are performed using pester’s “TestDrive” functionality

#### Notes

- This is how I wanted it! 
- It can use built-in functionality, but it is not clever enough to use it by it’s own devise.

## Key Take aways

- Dealing with complex subdomains that has each it’s owns do’s and dont’s into requirements (such as files, dates) adds on additional requirements and conciseness for the navigator how it is used.
- This conciseness and guidance caries over to tests even more as this is for some parts the main domain of the tests.
- The AI might be able to draw code like using TestDrive up a hat but only when navigated to. A very talented junior with no idea how to apply any of it’s knowledge directly.