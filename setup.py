from setuptools import setup, find_packages


def setup_package():
    setup(name='j',
          version='1.0.0',
          description='jira ticket issuer',
          author='Yuki Kanai',
          author_email='yukikanai0204@gmail.com',
          install_requires=['jira'],
          packages=find_packages(),
          entry_points={
              'console_scripts': [
                  'j = j.j:main'
              ]
          }
          )


if __name__ == '__main__':
    setup_package()
