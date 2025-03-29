function New-DutyRoster {
    <#
    .SYNOPSIS
    Generates a duty roster with randomly assigned names to dates.

    .DESCRIPTION
    This cmdlet generates a duty roster by randomly drawing names from the provided list and pairing them with dates.
    The dates are generated as consecutive Mondays starting from a given start date. If the list of names is exhausted,
    it will be reset to the full list of names. The start date is optional and defaults to the current date.

    .PARAMETER Names
    The list of names to be assigned to dates.

    .PARAMETER Count
    The number of duty assignments to generate.

    .PARAMETER StartDate
    The starting date for the duty roster. The default is the current date.

    .EXAMPLE
    New-DutyRoster -Names "Alice", "Bob", "Charlie" -Count 10

    .EXAMPLE
    New-DutyRoster -Names "Alice", "Bob", "Charlie" -Count 10 -StartDate "2025-03-25"

    #>

    [CmdletBinding()]
    param (
        [Parameter(Mandatory = $true)]
        [string[]]$Names,

        [Parameter(Mandatory = $true)]
        [int]$Count,

        [Parameter(Mandatory = $false)]
        [datetime]$StartDate = (Get-Date)
    )

    # Determine the first Monday from the StartDate
    $CurrentDate = $StartDate
    while ($CurrentDate.DayOfWeek -ne [System.DayOfWeek]::Monday) {
        $CurrentDate = $CurrentDate.AddDays(1)
    }

    $Roster = @()
    $RemainingNames = $Names.Clone()
    $Random = New-Object System.Random

    for ($i = 0; $i -lt $Count; $i++) {
        # Pick a random name from the remaining names
        $Index = $Random.Next(0, $RemainingNames.Count)
        $SelectedName = $RemainingNames[$Index]

        # Add the name and date to the roster
        $Roster += [pscustomobject]@{
            Date = $CurrentDate.ToString("yyyy-MM-dd")
            Name = $SelectedName
        }

        # Remove the selected name from the remaining names
        $RemainingNames = $RemainingNames | Where-Object { $_ -ne $SelectedName }

        # Reset the remaining names if the list is empty
        if ($RemainingNames.Count -eq 0) {
            $RemainingNames = $Names.Clone()
        }

        # Move to the next Monday
        $CurrentDate = $CurrentDate.AddDays(7)
    }

    return $Roster
}