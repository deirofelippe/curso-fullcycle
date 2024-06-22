#!/bin/bash

kafka-topics --bootstrap-server localhost:9092 --topic balances --create ;
kafka-topics --bootstrap-server localhost:9092 --topic transactions --create