import os
import subprocess
import tempfile
from itertools import combinations


def decrypt(encrypted_solution_path, passwords):

    tempdir = tempfile.mkdtemp()

    destination = os.path.join(tempdir, 'solution.plain')

    for p in passwords:
        print 'testing password: %s' % p
        try:
            r = subprocess.check_call(['openssl', 'aes-256-cbc', '-d', '-a',
                                       '-in', encrypted_solution_path,
                                       '-out', destination,
                                       '-pass', 'pass:' + p])
            if r == 0:
                print 'success'
                return p
        except subprocess.CalledProcessError:
            pass
        print 'fail'

    return None


path = '/Users/benmeier/Desktop/Hackfu/hackfu2016/container/Challenge 4/solution.txt.enc'

components = ['BRINGLIFETOTHEWASTELAND', 'proceedtoeden']
components = [c.lower() for c in components] + [c.upper() for c in components]
components.append('')

passwords = [''.join(c) for c in combinations(components, 2)]

decrypt(path, passwords)
