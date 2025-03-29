# Filename: New-DutyRoster.Tests.ps1

$here = Split-Path -Parent $MyInvocation.MyCommand.Path
Import-Module "$here\New-DutyRoster.psm1"

Describe 'New-DutyRoster' {
    Mock Get-Date { return [datetime]::Parse("2025-03-25 20:31:29") }

    Mock -CommandName 'Get-Random' -MockWith {
        param (
            [Parameter(Mandatory)]
            [int]$Minimum = 0,
            [Parameter(Mandatory)]
            [int]$Maximum
        )
        # Return a deterministic sequence of random numbers for testing
        $script:randomIndex = $script:randomIndex + 1
        return $script:randomIndex % $Maximum
    } -ParameterFilter {
        $Maximum -eq 3
    }

    BeforeAll {
        $script:randomIndex = -1
    }

    It 'Generates a duty roster with given inputs' {
        $Names = @("Alice", "Bob", "Charlie")
        $Count = 10

        $expectedOutput = @(
            [pscustomobject]@{ Date = "2025-03-31"; Name = "Alice" },
            [pscustomobject]@{ Date = "2025-04-07"; Name = "Bob" },
            [pscustomobject]@{ Date = "2025-04-14"; Name = "Charlie" },
            [pscustomobject]@{ Date = "2025-04-21"; Name = "Alice" },
            [pscustomobject]@{ Date = "2025-04-28"; Name = "Bob" },
            [pscustomobject]@{ Date = "2025-05-05"; Name = "Charlie" },
            [pscustomobject]@{ Date = "2025-05-12"; Name = "Alice" },
            [pscustomobject]@{ Date = "2025-05-19"; Name = "Bob" },
            [pscustomobject]@{ Date = "2025-05-26"; Name = "Charlie" },
            [pscustomobject]@{ Date = "2025-06-02"; Name = "Alice" }
        )

        $result = New-DutyRoster -Names $Names -Count $Count

        $result | Should -BeExactly $expectedOutput
    }
}