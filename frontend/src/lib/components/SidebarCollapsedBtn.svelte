<script lang="ts">
    export let ariaLabel: string;
    export let shortcut: string;
    export let icon: string;

    const kbd = shortcut.split('+');
</script>

<button
    class="
        w-[1.9rem] h-[1.9rem]
        flex items-center justify-center
        border-0 cursor-pointer text-secondary
        rounded-lg p-[.3rem] relative
        hover:bg-glass-light
        transition-all duration-200 ease-linear
        group"
    aria-label="{ariaLabel}"
>
    <div class="iconify {icon} w-full h-full text-secondary"></div>

    <span class="
        bg-glass-white absolute left-(--popup-spacing)
        border-1 border-glass-border rounded-lg
        flex items-center justify-center backdrop-blur-2xl
        gap-2 text-secondary px-[.6rem] py-[.4rem]
        transition-all duration-200 ease-linear
        hidden group-hover:flex">
        {#each kbd as k}
            <kbd class="border-1 border-glass-border rounded-sm text-xs leading-none py-[.2rem] px-[.4rem]">{k}</kbd>
        {/each}
    </span>
</button>
