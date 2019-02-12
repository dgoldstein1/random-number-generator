# random-number-generator
simple service generating a random number with prometheus metrics

## Run it

```
docker run -p 8080:8080 dgoldstein1/random-number-generator
```

## API

`/metrics` -- shows prometheus metrics for the service

`/randomNumber?max={int64}` -- generates a random number with a max


## Authors

* **David Goldstein** - [DavidCharlesGoldstein.com](http://www.davidcharlesgoldstein.com/?github-password-service) - [Decipher Technology Studios](http://deciphernow.com/)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
