# Filename: Get-OnDutyPerson.Tests.ps1

# Import the module or script containing the Get-OnDutyPerson cmdlet
. "$PSScriptRoot\Get-OnDutyPerson.ps1"

Describe "Get-OnDutyPerson" {

    # Mock the current date to ensure tests are independent of the actual current date
    Mock Get-Date { return [datetime]::Parse("2025-03-29") }

    # Test case: Valid duty roster file and default current date
    It "Should return the person on duty for the last passed date" {
        # Create mock duty roster data
        $mockDutyRosterFilePath = Join-Path -Path TestDrive -ChildPath "mockDutyRoster.psd1"
        $mockDutyRosterContent = @"
@{
    DutyRoster = @(
        @{ Date = "2025-03-27"; Name = "Alice" },
        @{ Date = "2025-03-28"; Name = "Bob" },
        @{ Date = "2025-03-29"; Name = "Charlie" }
    )
}
"@
        # Write the mock duty roster content to the file
        $mockDutyRosterContent | Out-File -FilePath $mockDutyRosterFilePath -Force

        $result = Get-OnDutyPerson -RosterFilePath $mockDutyRosterFilePath
        $result | Should -Be "On duty: Charlie"
    }

    # Test case: Valid duty roster file and a specific date
    It "Should return the person on duty for the specific date" {
        # Create mock duty roster data
        $mockDutyRosterFilePath = Join-Path -Path TestDrive -ChildPath "mockDutyRoster.psd1"
        $mockDutyRosterContent = @"
@{
    DutyRoster = @(
        @{ Date = "2025-03-27"; Name = "Alice" },
        @{ Date = "2025-03-28"; Name = "Bob" },
        @{ Date = "2025-03-29"; Name = "Charlie" }
    )
}
"@
        # Write the mock duty roster content to the file
        $mockDutyRosterContent | Out-File -FilePath $mockDutyRosterFilePath -Force

        $specificDate = [datetime]::Parse("2025-03-28")
        $result = Get-OnDutyPerson -RosterFilePath $mockDutyRosterFilePath -CurrentDate $specificDate
        $result | Should -Be "On duty: Bob"
    }

    # Test case: Valid duty roster file and a date with no passed entries
    It "Should return no duty entry found if no dates have passed" {
        # Create mock duty roster data
        $mockDutyRosterFilePath = Join-Path -Path TestDrive -ChildPath "mockDutyRoster.psd1"
        $mockDutyRosterContent = @"
@{
    DutyRoster = @(
        @{ Date = "2025-03-27"; Name = "Alice" },
        @{ Date = "2025-03-28"; Name = "Bob" },
        @{ Date = "2025-03-29"; Name = "Charlie" }
    )
}
"@
        # Write the mock duty roster content to the file
        $mockDutyRosterContent | Out-File -FilePath $mockDutyRosterFilePath -Force

        $futureDate = [datetime]::Parse("2025-03-26")
        $result = Get-OnDutyPerson -RosterFilePath $mockDutyRosterFilePath -CurrentDate $futureDate
        $result | Should -Be "No duty entry found for the specified date."
    }

    # Test case: Missing duty roster file
    It "Should return an error if the duty roster file is missing" {
        { Get-OnDutyPerson -RosterFilePath (Join-Path -Path TestDrive -ChildPath "nonexistent.psd1") } | Should -Throw -ErrorId "Unable to load the duty roster from the specified file path"
    }

    # Test case: Empty duty roster file
    It "Should return an error if the duty roster is empty" {
        # Create an empty duty roster file
        $emptyDutyRosterFilePath = Join-Path -Path TestDrive -ChildPath "emptyDutyRoster.psd1"
        @{} | Out-File -FilePath $emptyDutyRosterFilePath -Force

        { Get-OnDutyPerson -RosterFilePath $emptyDutyRosterFilePath } | Should -Throw -ErrorId "No duty roster found in the specified file"
    }
}