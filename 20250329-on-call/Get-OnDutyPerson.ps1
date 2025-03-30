function Get-OnDutyPerson {
    [CmdletBinding()]
    param (
        # Path to the PowerShell data file (.psd1) containing the duty roster.
        [Parameter(Mandatory = $true, Position = 0)]
        [string]$RosterFilePath,

        # The date to determine who is on duty. Defaults to the current date.
        [Parameter(Mandatory = $false)]
        [datetime]$CurrentDate = (Get-Date)
    )

    <#
    .SYNOPSIS
    Outputs the name of the person who is on duty based on the duty roster.

    .DESCRIPTION
    This cmdlet reads a PowerShell data file (.psd1) containing the duty roster specified as a list of hashtables with Date and Name keys.
    The person on duty is determined by what Date has last passed, given by a script parameter that defaults to the current date.

    .PARAMETER RosterFilePath
    The path to the PowerShell data file (.psd1) containing the duty roster.

    .PARAMETER CurrentDate
    The date to determine who is on duty. Defaults to the current date.

    .EXAMPLE
    Get-OnDutyPerson -RosterFilePath "C:\dutyRoster.psd1"

    .EXAMPLE
    Get-OnDutyPerson -RosterFilePath "C:\dutyRoster.psd1" -CurrentDate "2025-03-28"

    .NOTES
    This cmdlet requires the duty roster file to be a PowerShell data file (.psd1) that defines a list of hashtables with Date and Name keys.
    #>

    try {
        # Import the duty roster from the .psd1 file
        $DutyRoster = Import-PowerShellDataFile -Path $RosterFilePath
    } catch {
        Write-Error "Unable to load the duty roster from the specified file path: $RosterFilePath"
        return
    }

    if (-not $DutyRoster) {
        Write-Error "No duty roster found in the specified file: $RosterFilePath"
        return
    }

    # Convert the current date to a string in the format yyyy-MM-dd
    $currentDateString = $CurrentDate.ToString('yyyy-MM-dd')

    # Find the last date that has passed
    $lastDutyEntry = $DutyRoster.DutyRoster |
        Where-Object { $_.Date -le $currentDateString } |
        Sort-Object { [datetime]::ParseExact($_.Date, 'yyyy-MM-dd', $null) } -Descending |
        Select-Object -First 1

    if ($lastDutyEntry) {
        Write-Output "On duty: $($lastDutyEntry.Name)"
    } else {
        Write-Output "No duty entry found for the specified date."
    }
}

# Example duty roster data format (dutyRoster.psd1)
# @{
#     DutyRoster = @(
#         @{ Date = "2025-03-27"; Name = "Alice" },
#         @{ Date = "2025-03-28"; Name = "Bob" },
#         @{ Date = "2025-03-29"; Name = "Charlie" }
#     )
# }