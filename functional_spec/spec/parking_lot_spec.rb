require 'spec_helper'

RSpec.describe 'Parking Lot' do
  let(:pty) { PTY.spawn('parking_lot') }

  before(:each) do
    run_command(pty, "create_parking_lot 3\n")
  end

  it "can create a parking lot", :sample => true do
    expect(fetch_stdout(pty)).to end_with("Created a parking lot with 3 slots\n")
  end

  it "can park a car" do
    run_command(pty, "park KA-01-HH-3141 Black\n")
    expect(fetch_stdout(pty)).to end_with("Allocated slot number: 1\n")
  end

  it "can unpark a car" do
    run_command(pty, "park KA-01-HH-3141 Black\n")
    run_command(pty, "leave 1\n")
    expect(fetch_stdout(pty)).to end_with("Slot number 1 is free\n")
  end

  it "can report status" do
    run_command(pty, "park KA-01-HH-1234 White\n")
    run_command(pty, "park KA-01-HH-3141 Black\n")
    run_command(pty, "park KA-01-HH-9999 White\n")
    run_command(pty, "status\n")
    expect(fetch_stdout(pty)).to end_with(<<-EOTXT
Slot No.    Registration No    Colour
1           KA-01-HH-1234      White
2           KA-01-HH-3141      Black
3           KA-01-HH-9999      White
EOTXT
)
  end

  it "can find a car by registration" do
    run_command(pty, "park KA-01-HH-1234 White\n")
    run_command(pty, "park KA-01-HH-3141 Black\n")
    run_command(pty, "park KA-01-HH-9999 White\n")
    run_command(pty, "slot_number_for_registration_number BAD-NUMBER-0\n")
    expect(fetch_stdout(pty)).to end_with("Not found\n")
    run_command(pty, "slot_number_for_registration_number KA-01-HH-3141\n")
    expect(fetch_stdout(pty)).to end_with("2\n")
  end

  it "assigns slots closest to the entrance" do
    run_command(pty, "park KA-01-HH-1234 White\n")
    run_command(pty, "park KA-01-HH-3141 Black\n")
    run_command(pty, "park KA-01-HH-9999 White\n")
    run_command(pty, "leave 2\n")
    run_command(pty, "park KA-01-HH-3141 Black\n")
    expect(fetch_stdout(pty)).to end_with("Allocated slot number: 2\n")
  end

  it "will not assign when parking lot is full" do
    run_command(pty, "park KA-01-HH-1234 White\n")
    run_command(pty, "park KA-01-HH-3141 Black\n")
    run_command(pty, "park KA-01-HH-9999 White\n")
    expect(fetch_stdout(pty)).to end_with("Allocated slot number: 3\n")
    run_command(pty, "park N-4 Green\n")
    expect(fetch_stdout(pty)).to end_with("Sorry, parking lot is full\n")
  end

  it "can find cars by colour" do
    run_command(pty, "park KA-01-HH-1234 White\n")
    run_command(pty, "park KA-01-HH-3141 Black\n")
    run_command(pty, "park KA-01-HH-9999 White\n")
    run_command(pty, "slot_numbers_for_cars_with_colour White\n")
    expect(fetch_stdout(pty)).to end_with("1, 3\n")
  end

  it "fails when no cars of a given color are found" do
    run_command(pty, "park KA-01-HH-1234 White\n")
    run_command(pty, "park KA-01-HH-3141 Black\n")
    run_command(pty, "park KA-01-HH-9999 White\n")
    run_command(pty, "slot_numbers_for_cars_with_colour Green\n")
    expect(fetch_stdout(pty)).to end_with("Not found\n")
  end

  it "removes cars from the color index after they leave" do
    run_command(pty, "park KA-01-HH-1234 White\n")
    run_command(pty, "park KA-01-HH-3141 Black\n")
    run_command(pty, "slot_numbers_for_cars_with_colour White\n")
    expect(fetch_stdout(pty)).to end_with("1\n")
    run_command(pty, "leave 1\n")
    run_command(pty, "slot_numbers_for_cars_with_colour White\n")
    expect(fetch_stdout(pty)).to end_with("Not found\n")
  end

  it "removes cars from the registration index after they leave" do
    run_command(pty, "park KA-01-HH-1234 White\n")
    run_command(pty, "park KA-01-HH-3141 Black\n")
    run_command(pty, "slot_number_for_registration_number KA-01-HH-3141\n")
    expect(fetch_stdout(pty)).to end_with("2\n")
    run_command(pty, "leave 2\n")
    run_command(pty, "slot_number_for_registration_number KA-01-HH-3141\n")
    expect(fetch_stdout(pty)).to end_with("Not found\n")
  end

  it "does not unpark from an empty slot" do
    run_command(pty, "park KA-01-HH-1234 White\n")
    run_command(pty, "leave 2\n")
    expect(fetch_stdout(pty)).to end_with("Not found\n")
  end
  # pending "add more specs as needed"
end
