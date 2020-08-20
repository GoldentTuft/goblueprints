go build

cd lib

set d=../../

go build %d%synonyms
go build %d%available
go build %d%sprinkle
go build %d%coolify
go build %d%domainify

cd ../