from revivalkit import revive

# will be stored into 02_more_usages.py.coffin
# when file doesn't exist, ex_1 = list()
ex_1 = revive(list)

# will be stored into ex_2.coffin
ex_2 = revive(list, 'ex_2')

# will be stored into /tmp/ex_3.pickle
ex_3 = revive(list, '/tmp/ex_3.pickle')

var partnerLocationId int64 = 1
var partnerLocationId int64 = 1