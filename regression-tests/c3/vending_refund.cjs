const assert = require('node:assert/strict');
const path = require('node:path');
const { createJiti } = require(process.env.JITI_MODULE);
const load = createJiti(__filename, { fsCache: false, interopDefault: false });
const base = path.resolve(__dirname, '../../solutions/typescript/src/VendingMachine');
const Machine = load(path.join(base, 'VendingMachine.ts')).default;
const { Coin } = load(path.join(base, 'Coin.ts'));
const machine = Machine.getInstance();
const output = [];
const log = console.log;
console.log = (...args) => output.push(args.join(' '));
try {
  machine.addItem('A', 'Snack', 100, 2);
  machine.selectItem('A');
  machine.insertCoin(Coin.QUARTER);
  // The machine publishes this registered state through its public getter.
  machine.getItemSelectedState().refund();
  assert(output.includes('Refunding: 25'));
  assert.equal(machine.getBalance(), 0);
  machine.selectItem('A');
  machine.insertCoin(Coin.QUARTER);
  assert.equal(machine.getBalance(), 25);
  machine.getItemSelectedState().refund();
} finally { console.log = log; }
console.log('PASS partial-payment refund through machine-owned state');
