from flask import Flask, Response
import os

app = Flask(__name__)

def convert_number(value, input_format, output_format):
    # Convert input to integer
    if input_format == 'dec':
        num = int(value, 10)
    elif input_format == 'bin':
        num = int(value, 2)
    elif input_format == 'hex':
        num = int(value, 16)
    else:
        raise ValueError(f"Invalid input format: {input_format}")

    # Convert integer to output format
    if output_format == 'dec':
        return str(num)
    elif output_format == 'bin':
        return bin(num)[2:]  # Remove '0b' prefix
    elif output_format == 'hex':
        return hex(num)[2:]  # Remove '0x' prefix
    else:
        raise ValueError(f"Invalid output format: {output_format}")

@app.route('/convert/<value>/<input_format>/<output_format>', methods=['GET'])
def convert(value, input_format, output_format):
    try:
        result = convert_number(value, input_format, output_format)
        return Response(result, mimetype='text/plain')
    except ValueError as e:
        return Response(str(e), status=400, mimetype='text/plain')

@app.route('/health', methods=['GET'])
def health_check():
    return Response("OK", mimetype='text/plain')

@app.route('/', defaults={'path': ''})
@app.route('/<path:path>')
def catch_all(path):
    usage_guide = """
Usage Guide:
GET /convert/<value>/<input-format>/<output-format>

<value>: Any alphanumeric value in the format specified by <input-format>
<input-format> and <output-format>:
  dec: Decimal (base-10) format
  bin: Binary (base-2) format
  hex: Hexadecimal (base-16) format

Example: /convert/1010/bin/dec
    """
    return Response(usage_guide, mimetype='text/plain')

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=8080)
