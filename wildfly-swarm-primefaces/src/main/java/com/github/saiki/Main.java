package com.github.saiki;

import org.wildfly.swarm.Swarm;
import org.wildfly.swarm.jsf.JSFFraction;

public class Main {

	public static void main(String[] args) throws Exception {
		Swarm swarm = new Swarm();
		swarm.fraction(new JSFFraction());
		swarm.start();
		swarm.deploy();
	}
}