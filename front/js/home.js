function checkinput(id_input) {

    $(document).ready(function() {
        input = $(id_input).val();

        if (input <= 8) {
            return false;
        }
        return true
    });
}

function printAlert(id_input, id_alert) {
    $(document).ready(function() {
        alert = $(id_input)
        checkinput = checkinput(id_input)

        if (!checkinput)
            alert.val = "the input is lower 8"
        else
            alert.val = ""

    });
}