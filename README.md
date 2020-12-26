# lab4 (variant 6)

### input(file : example.txt):
`print world`\
`print hello world`\
`pr a`\
`aaa`\
`split a:bc:d:ef :`\
`split as:s;q.gg`\
`spl` \
`sp a`

### output(console):
`world`\
`hello world`\
`PARSING ERROR: Invalid command: pr a`\
`PARSING ERROR: Invalid command: aaa`\
`SYNTAX ERROR: Invalid count of arguments for split: 2`\
`PARSING ERROR: Invalid command: spl`\
`PARSING ERROR: Invalid command: sp a`\
`a`\
`bc`\
`d`\
`ef`
