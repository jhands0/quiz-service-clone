<script lang="ts">
    import Leaderboard from "../../lib/Leaderboard.svelte";
    import { leaderboard } from "../../service/host/host";
</script>

<div class="flex justify-center bg-purple-500 min-h-screen w-full">
    <div class="mt-32">
        <h2 class="text-white text-center text-5xl font-bold"> Game ended! </h2>
        <div class="flex flex-wrap gap-2 mt-10">
            <Leaderboard leaderboard={$leaderboard} />
        </div>
    </div>
</div>