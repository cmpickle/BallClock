let BallClock = require('./ball-clock');

if (process.argv.length <= 1) {
  console.log(
    'Please enter a valid command line option:\n\t[--repetitionTime repetition-number]\n\t[--timedOutput number-of-balls amount-of-time]'
  );
  return;
}

switch (process.argv[2]) {
  case '--repetitionTime':
    balls = process.argv[3].toString();
    if (balls === undefined) {
      console.error('repetition time requires an integer parameter between 27 and 127.');
      return;
    }
    if (balls < 27 || balls > 127) {
      console.error('repetition time requires an integer parameter between 27 and 127.');
      return;
    }
    ballClock = new BallClock(balls);
    let days = 0;
    while (true) {
      ballClock.ticks(1440);
      days++;
      if (ballClock.IsStartingOrder()) {
        break;
      }
    }

    console.log(`${balls} balls cycle after ${days} days.`);
    break;
  case '--timedOutput':
    balls = process.argv[3].toString();
    if (balls === undefined) {
      console.error('repetition time requires an integer parameter balls between 27 and 127.');
      return;
    }
    if (balls < 27 || balls > 127) {
      console.error('repetition time requires an integer parameter balls between 27 and 127.');
      return;
    }
    let minutes = parseInt(process.argv[4]);
    if (minutes === undefined) {
      console.error('repetition time requires an integer parameter minutes greater than 0.');
      return;
    }
    if (minutes < 1) {
      console.error('repetition time requires an integer parameter minutes greater than 0.');
      return;
    }
    ballClock = new BallClock(balls);
    ballClock.ticks(minutes);
    console.log(JSON.stringify(ballClock));
    break;
  default:
    console.log(
      'Please enter a valid command line option:\n\t[--repetitionTime repetition-number]\n\t[--timedOutput number-of-balls amount-of-time]'
    );
    break;
}
