// use std::fmt;

#[derive(Debug, PartialEq)]
pub struct Clock {
    hours: i32,
    minutes: i32,
}

// impl PartialEq for Clock {
//     fn eq(&self, other: &Self) -> bool {
//         self.hours == other.hours && self.minutes == other.minutes
//     }
// }
// Clock::new(1, -40).to_string(), "00:20");

const MIDNIGHT_HOUR: i32 = 24;
const MINUTES_IN_HOUR: i32 = 60;
impl Clock {
    pub fn new(hours: i32, minutes: i32) -> Self {
        // Get the base minutes value
        let mut m = minutes % MINUTES_IN_HOUR;

        // Get additional hours from rollover minutes
        let hours_diff;

        // If minutes are below 0 then calculate hours_diff
        if m < 0 {
            // Add the negative value to minutes in hour to get subtract
            // from the revious hour.
            m = MINUTES_IN_HOUR + m;
            hours_diff = (minutes / MINUTES_IN_HOUR) - 1;
        } else {
            hours_diff = minutes / MINUTES_IN_HOUR
        }

        // Calculate hours
        let mut h = (hours + hours_diff) % MIDNIGHT_HOUR;

        // If negative hours then subtract from MIDNIGHT_HOUR
        if h < 0 {
            h = MIDNIGHT_HOUR + h;
        }

        Self {
            hours: h,
            minutes: m,
        }
    }

    pub fn add_minutes(&self, minutes: i32) -> Self {
        // unimplemented!("Add {} minutes to existing Clock time", minutes);
        // &self

        Self {
            hours: self.hours,
            minutes: self.minutes + minutes,
        }
    }
    pub fn to_string(&self) -> String {
        format!("{:0>2}:{:0>2}", self.hours, self.minutes)
    }
}
