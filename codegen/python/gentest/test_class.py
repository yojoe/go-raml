import unittest

import classtest.file as file
import classtest.dir as dir


class TestRAMLScalarAsTypeName(unittest.TestCase):
    '''
    Test using RAML scalar type as type name
    '''

    def testClassCreation(self):
        f = file(name='thefile')
        assert f is not None

        d = dir(files=[f])
        assert d is not None


if __name__ == '__main__':
    unittest.main()
