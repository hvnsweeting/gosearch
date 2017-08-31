#!/usr/bin/awk -f
# Download github.com/avelino/awesome-go data: `wget https://raw.githubusercontent.com/avelino/awesome-go/master/README.md`.
# Usage: ./gosearch.awk README.md <package_name> [category_filter]

BEGIN {
    FS="^\* \[|\]\(|\) - "

    if (length(ARGV[2]) == 0) {
	print "missing search term"
	exit 1
    }
    search_term=ARGV[2]

    if (length(ARGV[3]) != 0) {
	cat_filter=ARGV[3]
    }
}

/^## / {
    sub(/^## /, "")
    cat=$0
}

/^\* \[.+\]\(.+\) - (.+)$/{
    if (length(cat_filter) !=0 && cat != cat_filter) {
	next
    }
    if (index($0, search_term)) {
	print "---"
	sub(/^https?:\/\//, "", $3)
	print "Package: "$3
	print "Category: "cat
	print "Description: "$4
    }
}

/^# / {
    if (NR != 1) {
	exit
    }
}
