<#
.SYNOPSIS
Generates a duty roster by pairing names with sequential Mondays starting from a given date.

.DESCRIPTION
This cmdlet takes a list of names as an argument and produces a new list by pairing each name with a date. 
The dates are generated by taking the following Monday from the last computed date, starting with a date given 
as an argument. If no date is provided, the current date is used as the starting point.

.PARAMETER Names
A list of names to be paired with dates.

.PARAMETER StartDate
The starting date to begin generating the dates. If not provided, the current date is used.

.PARAMETER Count
The number of times a name and date pair will be selected.

.EXAMPLE
PS> New-DutyRoster -Names @("Alice", "Bob", "Charlie") -StartDate "2025-03-23" -Count 5

.EXAMPLE
PS> New-DutyRoster -Names @("Alice", "Bob", "Charlie") -Count 5

#>

function New-DutyRoster {
    [CmdletBinding()]
    param (
        [Parameter(Mandatory=$true)]
        [string[]]$Names,

        [Parameter(Mandatory=$false)]
        [datetime]$StartDate = (Get-Date),

        [Parameter(Mandatory=$true)]
        [int]$Count
    )

    function Get-NextMonday {
        param (
            [datetime]$date
        )
        $daysToAdd = (7 - $date.DayOfWeek + [DayOfWeek]::Monday) % 7
        if ($daysToAdd -eq 0) {
            $daysToAdd = 7
        }
        return $date.AddDays($daysToAdd)
    }

    $currentDate = Get-NextMonday -date $StartDate
    $result = @()
    $availableNames = $Names.Clone()

    for ($i = 0; $i -lt $Count; $i++) {
        if ($availableNames.Count -eq 0) {
            $availableNames = $Names.Clone()
        }

        $randomIndex = Get-Random -Maximum $availableNames.Count
        $selectedName = $availableNames[$randomIndex]
        $availableNames = $availableNames | Where-Object { $_ -ne $selectedName }

        $result += [PSCustomObject]@{
            Name = $selectedName
            Date = $currentDate.ToString("yyyy-MM-dd")
        }
        $currentDate = Get-NextMonday -date $currentDate
    }

    return $result
}
