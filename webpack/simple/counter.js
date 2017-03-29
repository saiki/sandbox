var Counter = function(init) {
	init = (init ? init : 0);
	this.count = init;
}

Counter.prototype.countUp = function() {
	return this.count++;
}

Counter.prototype.countDown = function() {
	return this.count--;
}

Counter.prototype.current = function() {
	return this.count;
}

module.exports = Counter;
