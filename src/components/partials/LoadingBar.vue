<script lang="ts" setup>
const props = withDefaults(
    defineProps<{
        throttle?: number;
        duration?: number;
        height?: number;
    }>(),
    {
        throttle: 150,
        duration: 2250,
        height: 3,
    },
);

// Options & Data
const data = reactive({
    percent: 0,
    left: 0,
    show: true,
    canSucceed: true,
});

// Local variables
let _timer: undefined | number | NodeJS.Timeout;
let _throttle: undefined | NodeJS.Timeout;
const _cut = 10000 / Math.floor(props.duration);

// Functions
const clear = () => {
    _timer && clearInterval(_timer);
    _throttle && clearTimeout(_throttle);
    _timer = undefined;
};
const start = () => {
    clear();
    data.percent = 0;

    if (props.throttle) {
        _throttle = setTimeout(startTimer, props.throttle);
    } else {
        startTimer();
    }
};
const increase = (num: number) => {
    data.percent = Math.min(100, Math.floor(data.percent + num));
};
const finish = () => {
    data.percent = 100;
    hide();
};
const hide = () => {
    clear();
    setTimeout(() => {
        data.show = false;
        setTimeout(() => {
            data.percent = 0;
        }, 550);
    }, 500);
};
const startTimer = () => {
    data.show = true;
    _timer = setInterval(() => {
        increase(_cut);
    }, 100);
};

const delayedFinish = () => {
    data.percent = 70;
    setTimeout(() => {
        finish();
    }, 500);
};

// Hooks
const nuxt = useNuxtApp();

// @ts-ignore we are currently unable to add custom event types to the typings
nuxt.hook('data:loading:start', start);
// @ts-ignore we are currently unable to add custom event types to the typings
nuxt.hook('data:loading:finish', delayedFinish);
// @ts-ignore we are currently unable to add custom event types to the typings
nuxt.hook('data:loading:finish_error', () => {
    data.canSucceed = false;
    delayedFinish();
    setTimeout(() => {
        data.canSucceed = true;
    }, 1250);
});

onBeforeUnmount(() => clear);
</script>

<template>
    <div
        :class="['nuxt-progress', !data.canSucceed ? 'nuxt-progress-failed' : '']"
        :style="{
            width: data.percent + '%',
            left: data.left,
            height: props.height + 'px',
            opacity: data.show ? 1 : 0,
            backgroundSize: (100 / data.percent) * 100 + '% auto',
        }"
    />
</template>

<style scoped>
.nuxt-progress {
    position: fixed;
    top: 0px;
    left: 0px;
    right: 0px;
    width: 0%;
    opacity: 1;
    transition:
        width 0.1s,
        height 0.4s,
        opacity 0.4s;
    background: repeating-linear-gradient(to right, #008668 0%, #00a77d 50%, #19dda5 100%);
    z-index: 999999;
}

.nuxt-progress-failed {
    background: repeating-linear-gradient(to right, #d72638 0%, #ac1e2d 50%, #d72638 100%);
}
</style>
