class Parser {
  constructor() {
    this.input = '';
    this.position = 0;
  }

  readLine() {
    return this.input.slice(this.position, this.position + 1);
  }

  readNumber() {
    return parseFloat(this.readLine());
  }

  readString() {
    return this.readLine();
  }

  readArray() {
    return this.readLine().split(' ');
  }

  readObject() {
    return JSON.parse(this.readLine());
  }

  readLineArray() {
    return this.readLine().split('\n');
  }

  readNumberArray() {
    return this.readLineArray().map(Number);
  }

  readStringArray() {
    return this.readLineArray().map(s => s.trim());
  }

  readObjectArray() {
    return this.readLineArray().map(s => JSON.parse(s));
  }

  readObjectMap() {
    const map = {};
    const lines = this.readLineArray();
    for (const line of lines) {
      const [key, value] = line.split('=');
      map[key.trim()] = value.trim();
    }
    return map;
  }

  readMap() {
    const map = {};
    const lines = this.readLineArray();
    for (const line of lines) {
      const [key, value] = line.split('=');
      map[key.trim()] = value.trim();
    }
    return map;
  }

  readMapObject() {
    const map = {};
    const lines = this.readLineArray();
    for (const line of lines) {
      const [key, value] = line.split('=');
      map[key.trim()] = JSON.parse(value.trim());
    }
    return map;
  }

  readMapObjectMap() {
    const map = {};
    const lines = this.readLineArray();
    for (const line of lines) {
      const [key, value] = line.split('=');
      map[key.trim()] = JSON.parse(value.trim());
    }
    return map;
  }

  write() {
    return this.input;
  }

  writeArray(array) {
    return this.input + array.join(' ');
  }

  writeObject(obj) {
    return JSON.stringify(obj);
  }

  writeMap(map) {
    return this.input + JSON.stringify(map);
  }

  writeMapObject(obj) {
    return JSON.stringify(obj);
  }

  writeMapObjectMap(map) {
    return JSON.stringify(map);
  }
}

export default Parser;