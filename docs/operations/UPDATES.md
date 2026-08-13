# Updates

## MVP
Use Debian packages (`.deb`) and an authenticated package repository/update flow because it minimizes custom infrastructure.

Requirements:
- signed packages/repository;
- no unattended breaking migration without recovery;
- clear reboot requirement;
- version/rollback compatibility documented.

## Future
Evaluate transactional/A-B image update or `systemd-sysupdate` only after the product is stable enough to justify the added image/partition complexity.

Update mechanism changes require an ADR.
