import Counter from './counter.js';

var App = function(id) {
	this.counter = new Counter(0);
	this.id = id;
};

App.prototype.up = function() {
	this.counter.countUp();
	var element = document.getElementById(this.id);
	element.innerHTML = "<p>count: "+this.counter.current()+"</p>";
}

App.prototype.down = function() {
	this.counter.countDown();
	var element = document.getElementById(this.id);
	element.innerHTML = "<p>count: "+this.counter.current()+"</p>";
}

var app = new App("out");
window.app = app;
