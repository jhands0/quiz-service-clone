<script lang="ts">
    import type { Player } from "../../model/quiz"

    export let player: Player;
</script>

<button on:click>
    <h3 class="text-3x1 text-white px-4 py-2 rounded-xl hover:line-through" style="background-color: rgb(33, 33, 33, 0.3);">
        {player.name}
    </h3>
</button>