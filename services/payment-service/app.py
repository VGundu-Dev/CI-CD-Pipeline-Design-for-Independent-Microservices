from flask import Flask
app = Flask(__name__)

@app.route('/')
def hello_world():
    return '''
    <div style="text-align: center; margin-top: 50px; font-family: sans-serif;">
      <h1 style="color: #9C27B0;">Payment Service</h1>
      <p>Status: 🟢 <strong>Online</strong></p>
      <p>Version: 1.0</p>
      <p>Technology: Python/Flask</p>
    </div>
    '''

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5001)
