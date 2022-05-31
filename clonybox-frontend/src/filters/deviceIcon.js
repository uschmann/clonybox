

export default function deviceIcon(value) {
    switch (value) {
        case "Computer":
            return 'mdi-desktop-classic';
        case "Speaker":
            return 'mdi-speaker-wireless';
    }

    return 'mdi-speaker-wireless';
}