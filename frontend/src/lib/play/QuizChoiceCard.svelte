<script lang="ts">
    export let color: string;
</script>

<div class="w-1/2 text-white font-bold flex items-center text-3x1 {color}">
    <slot />
</div>