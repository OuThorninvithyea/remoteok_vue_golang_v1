import { ref, onMounted } from 'vue'
import axios from 'axios'

export function useSidebarStats() {
    const statsData = ref([])
    const loading = ref(true)
    const error = ref(null)

    onMounted(async () => {
        try {
            const res = await axios.get('/api/postjob/sidebar')
            statsData.value = res.data.stats
        } catch (e) {
            error.value = e.message
        } finally {
            loading.value = false
        }
    })

    return { statsData, loading, error }
}
