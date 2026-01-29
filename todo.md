## 🧠 Phase 0: Clarity First (10–15 min)

- [x] Decide the project name (repo name fixed)
- [x] Write 2 lines describing what the app does
- [x] Decide:

  - events are **predefined** (hardcoded)
  - data lives **only in memory**

✔ If you can explain your project in 2 sentences, you’re ready.

---

## 🏗 Phase 1: Project Skeleton

- [x] Create a new Go project folder
- [x] Initialize Go module
- [x] Create a single entry file (`main`)
- [x] Make sure:

  - program runs
  - prints a welcome message
  - exits cleanly

✔ Goal: _“My CLI starts correctly.”_

---

## 📦 Phase 2: Data Models (Very Important)

- [x] Define an **Event** structure
- [x] Define a **Booking** structure
- [x] Decide what fields each will have
- [x] Create an in-memory list for:

  - events
  - bookings

✔ Goal: _“I know what data my program stores.”_

---

## 🧪 Phase 3: Seed Sample Data

- [x] Create 2–3 sample events
- [x] Assign:

  - IDs
  - names
  - dates
  - ticket counts

- [x] Print all events once to verify data

✔ Goal: _“My data looks correct when printed.”_

---

## 🧭 Phase 4: Main Menu Loop

- [ ] Show menu options:

  - view events
  - view event details
  - book tickets
  - view bookings
  - exit

- [ ] Take user input
- [ ] Keep showing menu until user exits

✔ Goal: _“The CLI doesn’t exit unexpectedly.”_

---

## 📋 Phase 5: View Events Feature

- [ ] Loop through all events
- [ ] Display:

  - ID
  - name
  - available tickets

- [ ] Format output cleanly

✔ Goal: _“User can see all events clearly.”_

---

## 🔍 Phase 6: View Event Details

- [ ] Ask user for event ID
- [ ] Find matching event
- [ ] Display full details
- [ ] Handle invalid ID gracefully

✔ Goal: _“Wrong input doesn’t crash the program.”_

---

## 🎟 Phase 7: Book Tickets (Core Feature)

- [ ] Ask for event ID
- [ ] Ask for number of tickets
- [ ] Validate:

  - event exists
  - ticket count > 0
  - tickets available

- [ ] Reduce available tickets
- [ ] Create a booking
- [ ] Generate a booking ID
- [ ] Show confirmation message

✔ Goal: _“Booking changes system state correctly.”_

---

## 📚 Phase 8: View My Bookings

- [ ] Check if bookings exist
- [ ] Loop through bookings
- [ ] Display:

  - booking ID
  - event name
  - ticket count

- [ ] Show friendly message if none exist

✔ Goal: _“User can verify their actions.”_

---

## 🛑 Phase 9: Exit Cleanly

- [ ] Print goodbye message
- [ ] Exit program intentionally
- [ ] No panic, no forced termination

✔ Goal: _“Program ends professionally.”_

---

## 🧹 Phase 10: Polish & Cleanup

- [ ] Rename confusing variables
- [ ] Remove unused code
- [ ] Make error messages friendly
- [ ] Keep output readable

✔ Goal: _“Code looks intentional, not rushed.”_

---

## 🧪 Phase 11: Manual Testing Checklist

Test these **yourself**:

- [ ] Invalid menu choice
- [ ] Invalid event ID
- [ ] Booking more tickets than available
- [ ] Booking zero or negative tickets
- [ ] Multiple bookings for same event

✔ Goal: _“Nothing breaks easily.”_

---

## 📄 Phase 12: GitHub Ready

- [ ] Write README:

  - project description
  - features
  - how to run

- [ ] Add sample output (optional)
- [ ] Push to GitHub

✔ Goal: _“Anyone can understand your project.”_

---

## 🧠 Mental Rule While Building

If you feel stuck:

> **Print everything and verify state.**

That’s how backend engineers debug.

---

## 🔜 After This (Only if you want upgrades)

Do **one** at a time:

- save data to file
- cancel booking
- seat numbers
- login system

---

If you want next, I can:

- turn this into a **daily build plan**
- review your checklist progress
- help you design **Version 2 features cleanly**

You picked a solid project. Build it slowly and cleanly — that’s how good engineers are made 🚀
