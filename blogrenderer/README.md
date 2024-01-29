
# Table of Contents

1.  [View post](#org0df3ffb)
    1.  [Tests](#org7509cd3)
2.  [Index](#org7d14d85)
    1.  [Tests](#orgf132802)
3.  [Design](#orge67533c)
4.  [html snippets](#org6adf11a)



<a id="org0df3ffb"></a>

# View post

-   renders a specific post
    -   the @Body field in @Post is a string containing markdown that should be converted to HTML


<a id="org7509cd3"></a>

## Tests

-   post as input, html as output
    -   need to unmarshall html to test if it was written?
    -   or test if the desired string is inside the html


<a id="org7d14d85"></a>

# Index

-   lists all of the posts, with hyperlinks to view the specific post


<a id="orgf132802"></a>

## Tests

-   same question, how would you test that


<a id="orge67533c"></a>

# Design

-   accepts an io.Writer
    -   can be os.Stdout, a fileWriter, a HTMLWriter
    -


<a id="org6adf11a"></a>

# html snippets

    <!-- <html lang="en"> -->
    <!--   <head> -->
    <!--     <meta charset="UTF-8" /> -->
    <!--     <meta name="viewport" content="width=device-width, initial-scale=1.0" /> -->
    <!--     <meta content="text/html; charset=utf-8" /> -->
    <!--     <link rel="stylesheet" href="styles.css" /> -->
    <!--     <title>{{.Title}}</title> -->
    <!--   </head> -->
    <!--   <body> -->
    <!--   </body> -->
    <!-- </html> -->

