<#
.SYNOPSIS
Pester tests for the New-DutyRoster cmdlet.

.DESCRIPTION
This file contains Pester tests to ensure the functionality of the New-DutyRoster cmdlet.

#>

# Import the module containing the New-DutyRoster cmdlet
Import-Module "$PSScriptRoot\New-DutyRoster.ps1"

# Describe block for New-DutyRoster cmdlet
Describe "New-DutyRoster Cmdlet Tests" {
    Context "When generating a duty roster" {
        Mock Get-Random { return 0 }

        It "Should generate the correct number of pairs" {
            $Names = @("Alice", "Bob", "Charlie")
            $Count = 5
            $result = New-DutyRoster -Names $Names -Count $Count

            $result | Should -HaveCount $Count
        }

        It "Should include all names before repeating any" {
            $Names = @("Alice", "Bob", "Charlie")
            $Count = 5
            $result = New-DutyRoster -Names $Names -Count $Count

            $uniqueNames = $result.Name | Select-Object -Unique
            $uniqueNames | Should -Contain "Alice"
            $uniqueNames | Should -Contain "Bob"
            $uniqueNames | Should -Contain "Charlie"
        }

        It "Should start from the given start date" {
            $Names = @("Alice", "Bob", "Charlie")
            $Count = 3
            $StartDate = [datetime]"2025-03-23"
            $result = New-DutyRoster -Names $Names -Count $Count -StartDate $StartDate

            $result[0].Date | Should -Be "2025-03-24"  # First Monday after 2025-03-23
            $result[1].Date | Should -Be "2025-03-31"  # Next Monday
            $result[2].Date | Should -Be "2025-04-07"  # Next Monday
        }

        It "Should default to the current date if no start date is provided" {
            Mock Get-Date { return [datetime]"2025-03-23" }

            $Names = @("Alice", "Bob", "Charlie")
            $Count = 3
            $currentDate = Get-Date
            $nextMonday = (Get-Date).AddDays((7 - (Get-Date).DayOfWeek + [DayOfWeek]::Monday) % 7)
            if ($nextMonday -eq $currentDate) {
                $nextMonday = $nextMonday.AddDays(7)
            }

            $result = New-DutyRoster -Names $Names -Count $Count

            $result[0].Date | Should -Be $nextMonday.ToString("yyyy-MM-dd")
        }

        It "Should produce the correct output for the given example" {
            Mock Get-Date { return [datetime]"2025-03-23" }
            Mock Get-Random { param ($Maximum); return 0 }

            $Names = @("Alice", "Bob", "Charlie")
            $StartDate = [datetime]"2025-03-23"
            $Count = 5

            $expectedResult = @(
                [PSCustomObject]@{ Name = "Alice"; Date = "2025-03-24" }
                [PSCustomObject]@{ Name = "Bob"; Date = "2025-03-31" }
                [PSCustomObject]@{ Name = "Charlie"; Date = "2025-04-07" }
                [PSCustomObject]@{ Name = "Alice"; Date = "2025-04-14" }
                [PSCustomObject]@{ Name = "Bob"; Date = "2025-04-21" }
            )

            $result = New-DutyRoster -Names $Names -StartDate $StartDate -Count $Count

            $result | Should -BeExactly $expectedResult
        }
    }
}