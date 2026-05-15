import { ref, onMounted } from 'vue'
import axios from 'axios'

export function useDropdownSearch() {
    const searchData = ref([])
    const loading = ref(true)
    const error = ref(null)

    onMounted(async () => {
        try {
            const res = await axios.get('/api/home/dropdown/search')
            searchData.value = res.data
        } catch (e) {
            error.value = e.message
        } finally {
            loading.value = false
        }
    })

    return { searchData, loading, error }
}
