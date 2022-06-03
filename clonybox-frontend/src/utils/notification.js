import EventBus from "./EventBus";

const notification = {
    info(text)
    {
        EventBus.$emit('alert', {text, icon: 'mdi-information-outline', color: 'info'});
    },
    success(text)
    {
        EventBus.$emit('alert', {text, icon: 'mdi-information-outline', color: 'success'});
    },
    error(text)
    {
        EventBus.$emit('alert', {text, icon: 'mdi-error_outline', color: 'error'});
    },
};

export default notification;
