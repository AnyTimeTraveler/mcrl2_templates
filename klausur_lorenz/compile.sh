set -x

mcrl22lps parkhaus.mcrl2 parkhaus.lps --lin-method=regular2

lps2lts parkhaus.lps parkhaus.lts

ltsconvert parkhaus.lts parkhaus.lts --equivalence=branching-bisim

ltsview parkhaus.lts --verbose
