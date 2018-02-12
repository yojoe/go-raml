import unittest

from goramldir.server.app import app
from goramldir.server.types.User import User


class TestBase(unittest.TestCase):

    def setUp(self):
        self.app = app.test_client()
        self.app.testing = True


class TestValidation(TestBase):

    def testValidInput(self):
        user = User(username='me')
        rv = self.app.post('/users', data=user.as_json(), content_type='application/json')
        self.assertEqual(rv.status_code, 200)

    def testInvalidInput(self):
        user = {'city': 'bandung'}
        rv = self.app.post('/users', data=user, content_type='application/json')
        self.assertEqual(rv.status_code, 400)


if __name__ == '__main__':
    unittest.main()
