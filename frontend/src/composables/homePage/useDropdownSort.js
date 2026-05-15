import { ref, onMounted } from 'vue'
import axios from 'axios'

export function useDropdownSort() {
    const sortData = ref([])
    const loading = ref(true)
    const error = ref(null)

    onMounted(async () => {
        try {
            const res = await axios.get('/api/home/dropdown/sort')
            sortData.value = res.data
        } catch (e) {
            error.value = e.message
        } finally {
            loading.value = false
        }
    })

    return { sortData, loading, error }
}
