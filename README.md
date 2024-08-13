# Snowflake ID Generator

This repository contains a Golang implementation of a simple and efficient 34-bit unique ID generator. It is inspired by Twitter's Snowflake algorithm but designed with a much smaller bit length, making it ideal for specific use cases that require compact unique identifiers.

## Features

- **Unique 34-bit ID Generation**: The generator creates a unique 34-bit identifier composed of a timestamp and sequence number.
- **Sequence Handling**: If the sequence is exhausted within the same millisecond, the generator waits until the next millisecond to continue generating IDs.

## ID Structure

The 34-bit identifier is structured as follows:

- **32 bits** are used for the timestamp, representing milliseconds since the custom epoch .
- **2 bits** are used for the sequence number, which increments for each ID generated within the same millisecond.
