import unittest
import json

from goramldir.server.app import app as flask_app
from goramldir.server.types.User import User

from goramldir.sanic.app import app as sanic_app


class TestBase(unittest.TestCase):

    def setUp(self):
        self.flask_app = flask_app.test_client()
        self.flask_app.testing = True

        self.sanic_app = sanic_app.test_client

    def flask_post(self, endpoint, data):
        return self.flask_app.post(endpoint, data=data, content_type='application/json')

    def sanic_post(self, endpoint, data):
        return self.sanic_app.post(endpoint, data=data)


class TestFlaskValidation(TestBase):

    def testSimpleValidInput(self):
        '''
        Test simple validation with valid input
        '''
        user = User(username='me')

        rv = self.flask_post('/users', user.as_json())
        self.assertEqual(rv.status_code, 200)

        req, resp = self.sanic_post('/users', user.as_json())
        assert resp.status == 200

    def testSimpleInvalidInput(self):
        '''
        test simple validation with invalid input
        '''
        user = {'city': 'bandung'}

        rv = self.flask_post('/users', user)
        assert rv.status_code == 400

        req, resp = self.sanic_post('/users', json.dumps(user))
        assert resp.status == 400


if __name__ == '__main__':
    unittest.main()
